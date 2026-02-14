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

        // Collect sections
        const sections = [];
        const sectionGroups = document.querySelectorAll('.section-group');
        console.log("Found section groups:", sectionGroups.length);

        sectionGroups.forEach((group, index) => {
            const titleInput = group.querySelector('.section-title');
            const contentInput = group.querySelector('.section-content');
            
            console.log(`Processing section ${index}:`, titleInput?.value, contentInput?.value);

            if (titleInput && contentInput) {
                sections.push({
                    id: 'sect-' + Date.now() + '-' + Math.floor(Math.random() * 10000),
                    title: titleInput.value,
                    content: contentInput.value,
                    order: index + 1
                });
            }
        });

        console.log("Sending payload:", JSON.stringify({ 
            title, level, description, photo_url, general_info, sections 
        }, null, 2));

        const token = localStorage.getItem('token');
        await fetch("/api/courses", {
            method: "POST",
            headers: { 
                "Content-Type": "application/json",
                "Authorization": "Bearer " + token
            },
            body: JSON.stringify({ 
                title, level, description, photo_url, general_info, sections 
            })
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
