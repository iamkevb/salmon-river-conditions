function createFlowChart() {
  const data = {
    labels: [{{range .Times}}"{{.}}",{{end}}],
    datasets: [
      {
        data: [{{range .Readings}}{{.}},{{end}}],
        fill: false,
        borderColor: "rgba(54, 162, 235, 1)",
        pointRadius: 0,
      },
    ],
  };

  const flowToolTipTitle = (toolTipItems) => {
    return toolTipItems[0].label 
  }

  const flowToolTipLabel = (toolTipItem) => {
    return `${toolTipItem.formattedValue} cfs`
  }

  const config = {
    type: "line",
    data: data,
    options: {
      scales: {
        y: {
            beginAtZero: true // Start y-axis at zero
        },
      },
      maintainAspectRatio: false,
        plugins: {
          legend: false,
          tooltip: {
            displayColors: false,
            callbacks: {
                title: flowToolTipTitle,
                label: flowToolTipLabel,
              },
          }
        }
    }    
  }

  const ctx = document.getElementById("flowChart");
  new Chart(ctx, config);
}

document.addEventListener("DOMContentLoaded", () => {
  createFlowChart()
});
