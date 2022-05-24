const form = document.querySelector("#submit-file-form");
const jobContainer = document.querySelector(".print-jobs");
const jobMessages = document.querySelector(".submit-file-form-messages");

fetch("/printer/jobs")
.then(response => {
    return response.json();
})
.then(json => {
    for (let job of json) {
        jobContainer.appendChild(createJobElement(job.id, job.filename));
    }
});

form.addEventListener("submit", (e) => {
    e.preventDefault();
    fetch(form.action, {
        method: "POST",
        body: new FormData(form)
    })
    .then(response => {
        return response.json();
    })
    .then(json => {
        jobMessages.innerHTML = "";
        const flashMessages = json.flashMessages ?? [];
        for (let m of flashMessages) {
            const p = document.createElement("p");
            p.classList.add("form-error-message")
            p.innerText = m;

            jobMessages.appendChild(p);
        }
        if (json.success == true) {
            const job = createJobElement(json.jobID, json.filename);
            jobContainer.appendChild(job);
        }
    })
});

function createJobElement(id, filename) {
    const job = document.createElement("div");
    job.classList.add("job");
    {
        const jobName = document.createElement("h2");
        jobName.innerText = filename;
        jobName.classList.add("job-name");
        job.appendChild(jobName);
    }
    {
        const cancelJobButton = document.createElement("button");
        cancelJobButton.innerText = "Cancel";
        cancelJobButton.classList.add("cancel-job-button");

        // add click handler
        cancelJobButton.addEventListener("click", () => {
            fetch(`/printer/jobs/${id}`, {
                method: "DELETE",
                credentials: "same-origin"
            })
            .then(response => {
                if (response.ok) {
                    job.remove()
                }
            });
        });

        job.appendChild(cancelJobButton);
    }

    return job;
}