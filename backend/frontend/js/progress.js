function updateProgress(course, theory=false, examples=false, testScore=0) {
    const token = localStorage.getItem("token");
    if (!token) {
        alert("Сначала войдите в систему!");
        return;
    }

    fetch("https://YOUR_RENDER_URL/api/progress/update", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": "Bearer " + token
        },
        body: JSON.stringify({ course, theory, examples, testScore })
    })
    .then(res => res.json())
    .then(data => alert(`Прогресс обновлён для ${course}`))
    .catch(err => console.error(err));
}
