const apiBaseUrl = '/.netlify'

fetch(`${apiBaseUrl}/items`)
  .then(result => result.json())
  .then(reports => {
    reports.forEach((report) => {
      var template = document.querySelector('#item');
      template.content.querySelector('.title').textContent = report.title;
      var clone = document.importNode(template.content, true);
      document.body.appendChild(clone);
    })
  })
