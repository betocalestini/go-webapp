
async function ajaxSearch(search) {
  let information = document.getElementById('information');
  information.innerHTML = "<p>Buscando ...</p>";

  const init = {
    method: 'POST',
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
    },
    body: search
  }
  return await fetch('/products', init).then((resposta) => {
    if (resposta.ok) {
      information.innerHTML = ""
      return resposta.text();
    }
    throw new Error('Não foi possível realizar a busca');
  }).then(data => {
    if (data[0] == "0") {
      let arrayData = data.split('_')
      information.innerHTML = arrayData[1]
      dataTable = document.getElementById('table-div')
      dataTable.innerHTML = ""
    } else {
      let dataTable = document.getElementById('table-div');
      dataTable.innerHTML = data;
      aplicaDataTable('data-table')
    }
  });
}

async function aplicaDataTable(table) {
  let tableDestination = await new DataTable(`#${table}`, {
    order: [
      [0, 'asc'],
      [1, 'asc'],
      [2, 'asc']

    ],
    language: {
      url: '/assets/js/pt-BR.json'
    },
    paging: true,
    searching: true,
    colReorder: true,
    fixedHeader: true,
    dom: 'Bfrtip',
    type: 'text',
    buttons: [
      'copy', 'csv', 'excel', 'pdf', 'print', 'pageLength'
    ],
    lengthMenu: [
      [10, 50, 250, -1],
      ['10 linhas', '50 linhas', '250 linhas', 'Todas']
    ],
  });
}

document.addEventListener('DOMContentLoaded', function () {
  aplicaDataTable('data-table')
})

var search = document.getElementById('search')
search.addEventListener('keyup', () => {
  // console.log(search.value)
  ajaxSearch(search.value)
})



// var ajaxSearch = (search) => {
//   $.ajax({
//     type: 'POST',
//     dataType: 'html',
//     url: '/',
//     beforeSend: () => {
//       $('#data-table').html('Buscando...');
//     },
//     data: { search: search },
//     success: (r) => {
//       $('#data-table').html(r);
//     },
//   });
// };

// $('#search').keyup(() => {
//   ajaxSearch($('#search').val());
// });