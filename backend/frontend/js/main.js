const API = "/api/courses";
let coursesData = [];
let filteredCourses = [];
let currentView = 'grid'; // Default view

async function loadCourses() {
  try {
      const res = await fetch(API);
      coursesData = await res.json();
      filteredCourses = coursesData; // Initialize filtered list
      renderCourses();
  } catch (error) {
      console.error("Failed to load courses:", error);
      document.getElementById("courses").innerHTML = "<p>Error loading courses.</p>";
  }
}

function handleSearch() {
    const query = document.getElementById('searchInput').value.toLowerCase();
    
    filteredCourses = coursesData.filter(c => {
        return (c.title && c.title.toLowerCase().includes(query)) ||
               (c.description && c.description.toLowerCase().includes(query)) ||
               (c.author_name && c.author_name.toLowerCase().includes(query)) ||
               (c.level && c.level.toLowerCase().includes(query));
    });

    renderCourses();
}

function switchView(view) {
    currentView = view;
    
    // Update button states
    const btnGrid = document.getElementById('btnGrid');
    const btnList = document.getElementById('btnList');
    
    if (btnGrid && btnList) {
        btnGrid.className = view === 'grid' ? 'btn-toggle active' : 'btn-toggle';
        btnList.className = view === 'list' ? 'btn-toggle active' : 'btn-toggle';
    }

    renderCourses();
}

function renderCourses() {
  const container = document.getElementById("courses");
  container.innerHTML = "";

  if (!filteredCourses || filteredCourses.length === 0) {
    container.innerHTML = "<p style='text-align:center; color:#666;'>No courses found.</p>";
    return;
  }

  if (currentView === 'grid') {
      renderGrid(container);
  } else {
      renderList(container);
  }
}

function getCourseId(c) {
    return (typeof c.id === 'object' && c.id !== null) ? (c.id.$oid || JSON.stringify(c.id)) : c.id;
}

function getStarRatingHtml(rating, count) {
    const stars = Math.round(rating || 0);
    let html = '<span style="color: #ffc107; font-size: 1.1em;">';
    for (let i = 1; i <= 5; i++) {
        html += i <= stars ? '★' : '☆';
    }
    html += `</span> <span style="font-size: 0.85em; color: #666;">(${count || 0})</span>`;
    return html;
}

function renderGrid(container) {
    const grid = document.createElement('div');
    grid.className = 'courses-grid';

    filteredCourses.forEach(c => {
        const card = document.createElement('div');
        card.className = 'course-card';
        
        const photoSrc = c.photo_url || 'https://via.placeholder.com/300x160?text=No+Image';
        const courseId = getCourseId(c);
        const ratingHtml = getStarRatingHtml(c.average_rating, c.rating_count);

        card.innerHTML = `
            <img src="${photoSrc}" class="card-img" alt="${c.title}">
            <div class="card-body">
                <h3 class="card-title">${c.title}</h3>
                <div class="card-meta">
                    <span class="level-badge">${c.level || 'General'}</span>
                    <span>By ${c.author_name || 'Unknown'}</span>
                </div>
                <div style="margin-bottom: 10px;">${ratingHtml}</div>
                <p style="color:#666; font-size:0.9em; flex-grow:1; margin-bottom:15px; display: -webkit-box; -webkit-line-clamp: 3; -webkit-box-orient: vertical; overflow: hidden;">
                    ${c.description || 'No description available.'}
                </p>
                <div class="card-footer">
                    <a href="course_detail.html?id=${courseId}" style="display:block; width:100%; text-align:center; background:#007bff; color:white; padding:10px; border-radius:6px; text-decoration:none; font-weight:500;">Start Learning</a>
                </div>
            </div>
        `;
        grid.appendChild(card);
    });
    container.appendChild(grid);
}

function renderList(container) {
  const table = document.createElement("table");
  table.className = "course-table";
  
  table.innerHTML = `
    <thead>
      <tr>
        <th>Photo</th>
        <th>Course Title</th>
        <th>Rating</th>
        <th>Difficulty Level</th>
        <th>Author</th>
        <th>Action</th>
      </tr>
    </thead>
    <tbody>
    </tbody>
  `;

  const tbody = table.querySelector("tbody");

  filteredCourses.forEach(c => {
    const row = document.createElement("tr");
    const photo = c.photo_url ? `<img src="${c.photo_url}" alt="${c.title}" style="width: 50px; height: 50px; object-fit: cover; border-radius: 4px;">` : '—';
    const courseId = getCourseId(c);
    const ratingHtml = getStarRatingHtml(c.average_rating, c.rating_count);

    row.innerHTML = `
      <td>${photo}</td>
      <td><strong>${c.title}</strong></td>
      <td>${ratingHtml}</td>
      <td><span class="level-badge">${c.level || '—'}</span></td>
      <td>${c.author_name || 'Unknown'}</td>
      <td><a href="course_detail.html?id=${courseId}" style="color: #007bff; text-decoration: none; font-weight: bold;">Details</a></td>
    `;
    tbody.appendChild(row);
  });

  container.appendChild(table);
}

// Initial load
loadCourses();
