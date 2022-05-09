
function ajaxSearch(search) {
  let dataTable = document.getElementById('data-table');
  dataTable.innerHTML = "<p>Buscando ...</p>";

  const init = {
    method: 'POST',
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
    },
    body: search
  }
  return fetch('/', init).then((resposta) => {
    if (resposta.ok) {
      // console.log("Resposta:", resposta.text())
      // let dataTable = document.getElementById('data-table');
      // dataTable.innerHTML = resposta.body
      return resposta.text();
    }
    throw new Error('Não foi possível realizar a busca');
  }).then(data => {
    let dataTable = document.getElementById('data-table');
    dataTable.innerHTML = data;
  });
}



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