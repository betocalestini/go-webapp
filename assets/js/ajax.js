

// function ajax(method, url, object) {
//    var xhr = new XMLHttpRequest();
//    xhr.responseType = 'text';

//    xhr.onreadystatechange = () => {
//       if (search.value.length > 0) {
//          var table = document.getElementById('div-table');
//          table.innerHTML = '<p> Buscando.... </p>';
//       }
//       if (xhr.readyState === 4) {
//          table.innerHTML = xhr.responseText;
//          console.log(table);
//       }
//    };
//    xhr.open(method, url, true);

//    xhr.send(object);
// }

// var search = document.getElementById('search');
// search.addEventListener('keyup', () => {
//    ajax('POST', '/', search.value);
//    console.log(search.value);
// });
