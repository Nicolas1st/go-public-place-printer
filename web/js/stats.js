fetch("/stats/")
.then((response) => {
    return response.json();
})
.then((json) => {
    const dailyUsage = json.dailyUsage.reverse();

    const labels = [];
    for (let i = dailyUsage.length-1; i >=  0; i--) {
        let d = new Date();
        d.setDate(d.getDate()-i)
        d = d.toLocaleString("ru", {
            day: "numeric",
            month: "numeric",
            year: "numeric",
        });
        labels.push(d);
    }

    const data = {
        labels: labels,
        datasets: [
            {
                label: "Расход Бумаги За Последний Месяц",
                backgroundColor: "rgb(254, 99, 131)",
                borderColor: "rgb(254, 98, 131)",
                data: dailyUsage,
            },
        ],
    };

    const config = {
        type: "line",
        data: data,
    };

    const myChart = new Chart(document.getElementById("bar-chart"), config);
});
