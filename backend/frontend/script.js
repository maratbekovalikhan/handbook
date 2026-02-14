const form = document.getElementById("courseForm");
const coursesDiv = document.getElementById("courses");

if (form) {
    form.addEventListener("submit", async (e) => {
        e.preventDefault();

        const title = document.getElementById("title").value;
        const level = document.getElementById("level").value;
        const description = document.getElementById("description").value;
        const photo_url = document.getElementById("photo_url").value;
        const general_info = document.getElementById("general_info").value;

        const token = localStorage.getItem('token');
        await fetch("/api/courses", {
            method: "POST",
            headers: { 
                "Content-Type": "application/json",
                "Authorization": "Bearer " + token
            },
            body: JSON.stringify({ title, level, description, photo_url, general_info })
        });

        window.location.href = "index.html";
    });
}

if (coursesDiv) {
    fetch("/api/courses")
        .then(res => res.json())
        .then(data => {
            data.forEach(course => {
                const div = document.createElement("div");
                div.innerHTML = `<h3>${course.title}</h3><p>${course.level}</p><hr>`;
                coursesDiv.appendChild(div);
            });
        });
}
