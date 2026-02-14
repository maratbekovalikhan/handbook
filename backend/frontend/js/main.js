const API = "/api/courses";

async function loadCourses() {
  const res = await fetch(API);
  const courses = await res.json();

  const container = document.getElementById("courses");
  container.innerHTML = "";

  if (!courses || courses.length === 0) {
    container.innerHTML = "<p>No courses added yet</p>";
    return;
  }

  const table = document.createElement("table");
  table.className = "course-table";
  
  table.innerHTML = `
    <thead>
      <tr>
        <th>Photo</th>
        <th>Course Title</th>
        <th>Difficulty Level</th>
        <th>Author</th>
        <th>Action</th>
      </tr>
    </thead>
    <tbody>
    </tbody>
  `;

  const tbody = table.querySelector("tbody");

  courses.forEach(c => {
    const row = document.createElement("tr");
    const photo = c.photo_url ? `<img src="${c.photo_url}" alt="${c.title}" style="width: 50px; height: 50px; object-fit: cover; border-radius: 4px;">` : '—';
    
    // MongoDB ObjectID might be serialized as a string or an object with $oid
    const courseId = (typeof c.id === 'object' && c.id !== null) ? (c.id.$oid || JSON.stringify(c.id)) : c.id;

    row.innerHTML = `
      <td>${photo}</td>
      <td><strong>${c.title}</strong></td>
      <td><span class="level-badge">${c.level || '—'}</span></td>
      <td>${c.author_name || 'Unknown'}</td>
      <td><a href="course_detail.html?id=${courseId}" style="color: #007bff; text-decoration: none; font-weight: bold;">Details</a></td>
    `;
    tbody.appendChild(row);
  });

  container.appendChild(table);
}

loadCourses();
