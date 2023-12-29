function createPressureChart() {
  const data = {
    labels: [{{range .Times}}"{{.}}",{{end}}],
    datasets: [
      {
        data: [{{range .Pressure}}{{.}},{{end}}],
        fill: false,
        borderColor: "rgba(54, 162, 235, 1)",
        pointRadius: 0,
      },
    ],
  };

  const precipationToolTipTitle = (toolTipItems) => {
    return toolTipItems[0].label 
  }

  const precipationToolTipLabel = (toolTipItem) => {
    console.log(toolTipItem)
    return `${toolTipItem.formattedValue} hPa`
  }

  const config = {
    type: "line",
    data: data,
    options: {
      maintainAspectRatio: false,
        plugins: {
          legend: false,
          tooltip: {
            displayColors: false,
            callbacks: {
                title: precipationToolTipTitle,
                label: precipationToolTipLabel,
              },
          }
        }
    }    
  }

  const ctx = document.getElementById("pressureChart");
  new Chart(ctx, config);
}

document.addEventListener("DOMContentLoaded", () => {
  createPressureChart()
});
