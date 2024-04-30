/* global bootstrap: false */
(() => {
  'use strict'
  const tooltipTriggerList = Array.from(document.querySelectorAll('[data-bs-toggle="tooltip"]'))
  tooltipTriggerList.forEach(tooltipTriggerEl => {
    new bootstrap.Tooltip(tooltipTriggerEl)
  })
})()


function selectContainer(container) {
  // Remover la clase 'selected' de todos los contenedores
  const containers = document.querySelectorAll('.custom-container');
  containers.forEach(c => c.classList.remove('selected'));
  // Agregar la clase 'selected' al contenedor seleccionado
  container.classList.add('selected');
}

