const form = document.getElementById("courseForm");
const coursesDiv = document.getElementById("courses");

if (form) {
    form.addEventListener("submit", async (e) => {
        e.preventDefault();

        const title = document.getElementById("title").value;
        const level = document.getElementById("level").value;

        await fetch("/api/courses", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ title, level })
        });

        alert("Course added!");
        window.location.href = "/";
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
