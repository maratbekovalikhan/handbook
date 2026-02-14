package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"handbook/config"
	"handbook/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateCourse godoc
// @Summary Create a new course
// @Description create a new course with sections
// @Tags courses
// @Accept  json
// @Produce  json
// @Security Bearer
// @Param course body models.Course true "Course content"
// @Success 200 {object} map[string]string
// @Router /courses [post]
func CreateCourse(w http.ResponseWriter, r *http.Request) {
	var course models.Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, "Invalid input JSON", 400)
		return
	}

	if course.Sections == nil {
		course.Sections = []models.Section{}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get user from context (set by AuthMiddleware)
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusInternalServerError)
		return
	}

	var user models.User
	err = config.DB.Collection("users").FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	course.AuthorID = objID
	course.AuthorName = user.Name

	_, err = config.DB.Collection("courses").InsertOne(ctx, course)
	if err != nil {
		http.Error(w, "Error saving course", 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Course created"})
}

// GetCourses godoc
// @Summary Get all courses
// @Description get list of all courses
// @Tags courses
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Course
// @Router /courses [get]
func GetCourses(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := config.DB.Collection("courses").Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Error fetching courses", 500)
		return
	}

	var courses []models.Course
	cursor.All(ctx, &courses)

	json.NewEncoder(w).Encode(courses)
}

// GetCourse godoc
// @Summary Get course by ID
// @Description get details of a single course
// @Tags courses
// @Accept  json
// @Produce  json
// @Param id query string true "Course ID"
// @Success 200 {object} models.Course
// @Router /course [get]
func GetCourse(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", 400)
		return
	}

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "Invalid course ID", 400)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var course models.Course
	err = config.DB.Collection("courses").FindOne(ctx, bson.M{"_id": id}).Decode(&course)
	if err != nil {
		http.Error(w, "Course not found", 404)
		return
	}

	json.NewEncoder(w).Encode(course)
}

// DeleteCourse godoc
// @Summary Delete a course
// @Description delete a course by ID (Admin or Author only)
// @Tags courses
// @Accept  json
// @Produce  json
// @Security Bearer
// @Param id query string true "Course ID"
// @Success 200 {object} map[string]string
// @Router /course [delete]
func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	// Get User ID from context
	userID := r.Context().Value("userID").(string)
	userObjID, _ := primitive.ObjectIDFromHex(userID)

	// Get Course ID from URL
	courseIDStr := r.URL.Query().Get("id")
	if courseIDStr == "" {
		http.Error(w, "Missing id", 400)
		return
	}
	courseObjID, _ := primitive.ObjectIDFromHex(courseIDStr)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 1. Fetch User to check Role
	var user models.User
	err := config.DB.Collection("users").FindOne(ctx, bson.M{"_id": userObjID}).Decode(&user)
	if err != nil {
		http.Error(w, "User not found", 401)
		return
	}

	// 2. Fetch Course to check Author
	var course models.Course
	err = config.DB.Collection("courses").FindOne(ctx, bson.M{"_id": courseObjID}).Decode(&course)
	if err != nil {
		http.Error(w, "Course not found", 404)
		return
	}

	// 3. Authorization Check
	isAuthor := course.AuthorID == userObjID
	isAdmin := user.Role == "admin"

	if !isAuthor && !isAdmin {
		http.Error(w, "Forbidden: You don't have permission to delete this course", 403)
		return
	}

	// 4. Delete Course
	_, err = config.DB.Collection("courses").DeleteOne(ctx, bson.M{"_id": courseObjID})
	if err != nil {
		http.Error(w, "Error deleting course", 500)
		return
	}

	// Optional: Delete related progress and ratings
	config.DB.Collection("progress").DeleteMany(ctx, bson.M{"course_id": courseObjID})
	config.DB.Collection("ratings").DeleteMany(ctx, bson.M{"course_id": courseObjID})

	json.NewEncoder(w).Encode(map[string]string{"message": "Course deleted successfully"})
}
