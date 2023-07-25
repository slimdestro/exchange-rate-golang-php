 <!-- 
    - index.php: serves frontend 
    - bootstrap: used this to make UI responsive as well as its JS to easily implement filter
	- Author: Mukul(https://github.com/slimdestro) | https://www.modcode.dev
-->
<html>
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>AKCS-TEST | Dashboard</title> 
  <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
  
</head>

<body>
  <div class="container mt-5">
    <h1 class="mb-4">AKCS-TEST | Dashboard</h1>
    <input type="text" id="filterInput" placeholder="Filter by currency">
    <table class="table table-striped" id="currencyTable">
      <thead>
        <tr>
          <th>Currency</th>
          <th>Rate</th>
          <th>Date</th>
          <th>Timestamp</th>
        </tr>
      </thead>
      <tbody>

      </tbody>
    </table>
    <ul class="pagination" id="pagination"></ul>
  </div>
 
  <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.min.js"></script>
<script> 
  const API_URL = 'https://akc_test.dev/frontend/api.php'; 
  async function fetchData() {
    try {
      const response = await fetch(API_URL);
      const data = await response.json();
      return data;
    } catch (error) {
      console.error('Error fetching data:', error);
      return [];
    }
  }

  document.addEventListener('DOMContentLoaded', async () => {
    const data = await fetchData();
    const itemsPerPage = 25;
    let currentPage = 1; 
    function renderTableData(data) {
      const tableBody = document.querySelector('#currencyTable tbody');
      tableBody.innerHTML = '';
      data.forEach((row) => {
        const tr = document.createElement('tr');
        tr.innerHTML = `
          <td>${row.currency}</td>
          <td>${row.rate}</td>
          <td>${row.date}</td>
          <td>${row.timestamp}</td>
          
        `;
        tableBody.appendChild(tr);
      });
    }
 
    function paginateData(data, itemsPerPage, currentPage) {
      const startIndex = (currentPage - 1) * itemsPerPage;
      const endIndex = startIndex + itemsPerPage;
      return data.slice(startIndex, endIndex);
    }
 
    function renderPagination(data, itemsPerPage, currentPage) {
      const totalPages = Math.ceil(data.length / itemsPerPage);
      const paginationElement = document.querySelector('#pagination');
      paginationElement.innerHTML = '';

      for (let i = 1; i <= totalPages; i++) {
        const li = document.createElement('li');
        const link = document.createElement('a');
        link.href = '#';
        link.classList.add('page-link');
        link.textContent = i;
        link.addEventListener('click', () => {
          renderTableData(paginateData(data, itemsPerPage, i));
          renderPagination(data, itemsPerPage, i);
        });

        if (i === currentPage) {
          li.classList.add('page-item', 'active');
        } else {
          li.classList.add('page-item');
        }

        li.appendChild(link);
        paginationElement.appendChild(li);
      }
    }
 
    function filterData(data, keyword) {
      return data.filter((row) =>
        row.currency.toLowerCase().includes(keyword.toLowerCase())
      );
    }
 
    renderTableData(paginateData(data, itemsPerPage, currentPage));
    renderPagination(data, itemsPerPage, currentPage);
 
    const filterInput = document.getElementById('filterInput');
    filterInput.addEventListener('input', () => {
      const filteredData = filterData(data, filterInput.value);
      currentPage = 1; 
      renderTableData(paginateData(filteredData, itemsPerPage, currentPage));
      renderPagination(filteredData, itemsPerPage, currentPage);
    });
  });
</script>

</body>

</html>
