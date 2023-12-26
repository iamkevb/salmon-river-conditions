function createPressureChart() {
  const data = {
    labels: [{{range .Times}}"{{.}}",{{end}}],
    datasets: [
      {
        label: "My First Dataset",
        data: [{{range .Pressure}}{{.}},{{end}}],
        fill: false,
        borderColor: "rgba(54, 162, 235, 1)",
        tension: 0.1,
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
  };

  const ctx = document.getElementById("pressureChart");
  new Chart(ctx, config);
}

document.addEventListener("DOMContentLoaded", () => {
  createPressureChart();
});
