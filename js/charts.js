function createPressureChart(id, data) {
  const pressureToolTipTitle = (toolTipItems) => {
    return toolTipItems[0].label
  }

  const pressureToolTipLabel = (toolTipItem) => {
    return `${toolTipItem.formattedValue} hPa`
  }

  const config = {
    type: "line",
    data: data,
    options: {
      interaction: {
        mode: "nearest",
        intersect: false,
        axis: "x",
      },
      maintainAspectRatio: false,
      plugins: {
        legend: false,
        tooltip: {
          displayColors: false,
          callbacks: {
            title: pressureToolTipTitle,
            label: pressureToolTipLabel,
          },
        },
      },
      scales: {
        x: {
          ticks: {
            callback: (value, index) => {
              return index % 2 === 0 ? airData.labels[index] : '';
            }
          } 
        }
      }
    },
  }

  const ctx = document.getElementById(id)
  new Chart(ctx, config)
}

function createFlowChart(id, data) {
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
      interaction: {
        mode: "nearest",
        intersect: false,
        axis: "x",
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
        },
      },
      scales: {
        x: {
          ticks: {
            callback: (value, index) => {
              return index % 4 === 0 ? flowData.labels[index] : '';
            }
          } 
        }
      }
    },
  }

  const ctx = document.getElementById(id)
  new Chart(ctx, config)
}
