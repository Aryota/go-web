document.addEventListener('DOMContentLoaded', function() {
  var form = document.querySelector('form');
  var numberInput = document.getElementById('number');
  var resultElement = document.getElementById('result');
  var resetButton = document.getElementById('resetButton');

  form.addEventListener('submit', function(e) {
      e.preventDefault(); // フォームのデフォルトの送信動作をキャンセル

      var number = numberInput.value;
      resultElement.innerText = 'Input number: ' + number;
  });

  resetButton.addEventListener('click', function() {
      numberInput.value = ''; // 数値の入力欄をリセット
      resultElement.innerText = ''; // 結果表示をクリア
  });
});
