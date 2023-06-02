document.addEventListener('DOMContentLoaded', function() {
  var form = document.querySelector('form');
  var numberInput = document.getElementById('number');
  var tableBody = document.querySelector('#result-table tbody');
  var resetButton = document.getElementById('resetButton');

  form.addEventListener('submit', function(e) {
    e.preventDefault(); // フォームのデフォルトの送信動作をキャンセル

    var number = numberInput.value;

    // テーブルに行を追加
    var row = tableBody.insertRow();

    var numberCell = row.insertCell();
    var firstCell = row.insertCell();
    var secondCell = row.insertCell();
    var thirdCell = row.insertCell();

    numberCell.textContent = number;
    firstCell.textContent = getRandomNumber();
    secondCell.textContent = getRandomNumber();
    thirdCell.textContent = getRandomNumber();

    numberInput.value = ''; // 数値の入力欄をリセット
  });

  resetButton.addEventListener('click', function() {
    while (tableBody.firstChild) {
      tableBody.firstChild.remove(); // テーブルの行を全て削除
    }
  });

  // 0からmaxまでのランダムな整数を生成する関数
  function getRandomNumber() {
    var max = 60;
    return Math.floor(Math.random() * (max + 1));
  }
});
