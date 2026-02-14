package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"handbook/config"
	"handbook/models"

	"github.com/jung-kurt/gofpdf"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GenerateCertificate godoc
// @Summary Generate course certificate
// @Description generate a PDF certificate for a finished course
// @Tags certificate
// @Accept  json
// @Produce  application/pdf
// @Security Bearer
// @Param course_id query string true "Course ID"
// @Success 200 {file} file
// @Router /certificate [get]
func GenerateCertificate(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", 400)
		return
	}

	courseIDStr := r.URL.Query().Get("course_id")
	if courseIDStr == "" {
		http.Error(w, "Missing course_id", 400)
		return
	}
	courseObjID, err := primitive.ObjectIDFromHex(courseIDStr)
	if err != nil {
		http.Error(w, "Invalid course ID", 400)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1. Verify Progress
	var progress models.Progress
	err = config.DB.Collection("progress").FindOne(ctx, bson.M{
		"user_id":   userObjID,
		"course_id": courseObjID,
	}).Decode(&progress)

	if err != nil || !progress.IsFinished {
		http.Error(w, "You haven't finished this course yet", http.StatusForbidden)
		return
	}

	// 2. Fetch Course and User details
	var course models.Course
	config.DB.Collection("courses").FindOne(ctx, bson.M{"_id": courseObjID}).Decode(&course)

	var user models.User
	config.DB.Collection("users").FindOne(ctx, bson.M{"_id": userObjID}).Decode(&user)

	// 3. Create PDF
	pdf := gofpdf.New("L", "mm", "A4", "") // Landscape orientation
	pdf.AddPage()

	// Set background color/border
	pdf.SetLineWidth(2)
	pdf.SetHeaderFuncMode(nil, false)
	pdf.Rect(10, 10, 277, 190, "D")

	// Header - Handbook
	pdf.SetFont("Arial", "B", 34)
	pdf.SetTextColor(0, 123, 255) // #007bff
	pdf.CellFormat(0, 30, "HANDBOOK", "", 1, "C", false, 0, "")

	// Subtitle
	pdf.SetFont("Arial", "I", 16)
	pdf.SetTextColor(100, 100, 100)
	pdf.CellFormat(0, 10, "Certificate of Completion", "", 1, "C", false, 0, "")

	pdf.Ln(20)

	// Main content
	pdf.SetFont("Arial", "", 18)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 10, "This is to certify that", "", 1, "C", false, 0, "")

	pdf.Ln(5)

	// User Name
	pdf.SetFont("Arial", "B", 26)
	pdf.CellFormat(0, 15, user.Name, "", 1, "C", false, 0, "")

	// User Email (smaller)
	pdf.SetFont("Arial", "", 12)
	pdf.SetTextColor(100, 100, 100)
	pdf.CellFormat(0, 10, fmt.Sprintf("(%s)", user.Email), "", 1, "C", false, 0, "")

	pdf.Ln(10)

	pdf.SetFont("Arial", "", 18)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 10, "has successfully completed the course", "", 1, "C", false, 0, "")

	pdf.Ln(5)

	// Course Title
	pdf.SetFont("Arial", "B", 22)
	pdf.SetTextColor(0, 123, 255)
	pdf.CellFormat(0, 15, fmt.Sprintf("\"%s\"", course.Title), "", 1, "C", false, 0, "")

	pdf.Ln(15)

	// Author and Date
	completionDate := progress.FinishedAt.Format("January 02, 2006")
	if progress.FinishedAt.IsZero() {
		completionDate = time.Now().Format("January 02, 2006")
	}

	pdf.SetFont("Arial", "", 14)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 10, fmt.Sprintf("Instructor: %s", course.AuthorName), "", 1, "C", false, 0, "")
	pdf.CellFormat(0, 10, fmt.Sprintf("Date: %s", completionDate), "", 1, "C", false, 0, "")

	// Footer decoration
	pdf.Ln(10)
	pdf.SetDrawColor(0, 123, 255)
	pdf.Line(100, 180, 197, 180)

	// Set headers and output
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=Certificate_%s.pdf", course.Title))

	err = pdf.Output(w)
	if err != nil {
		http.Error(w, "Could not generate PDF", 500)
	}
}
