document.addEventListener('DOMContentLoaded', function() {
  var form = document.querySelector('#numberForm');
  var numberInput = document.getElementById('number');
  var resultBody = document.getElementById('result-body');
  var resetButton = document.getElementById('resetButton');
  var tableHeader = document.getElementById('table-header');

  // テーブルヘッダーを非表示にする
  tableHeader.style.display = 'none';

  form.addEventListener('submit', function(e) {
    e.preventDefault(); // フォームのデフォルトの送信動作をキャンセル

    var number = numberInput.value;

    // テーブルに行を追加
    var row = document.createElement('tr');

    var numberCell = document.createElement('td');
    var firstCell = document.createElement('td');
    var secondCell = document.createElement('td');
    var thirdCell = document.createElement('td');

    numberCell.textContent = number;
    firstCell.textContent = getRandomNumber();
    secondCell.textContent = getRandomNumber();
    thirdCell.textContent = getRandomNumber();

    row.appendChild(numberCell);
    row.appendChild(firstCell);
    row.appendChild(secondCell);
    row.appendChild(thirdCell);

    resultBody.appendChild(row);

    // テーブルヘッダーを表示
    tableHeader.style.display = 'table-row-group';

    numberInput.value = ''; // 数値の入力欄をリセット
  });

  resetButton.addEventListener('click', function() {
    resultBody.innerHTML = '';
    // テーブルヘッダーを非表示
    tableHeader.style.display = 'none';
  });

  // 0からmaxまでのランダムな整数を生成する関数
  function getRandomNumber() {
    var max = 60;
    return Math.floor(Math.random() * (max + 1));
  }
});
