{{define "form"}}
<!DOCTYPE html>
<html>
<body>
    <form action="./" method="get">
        <input type="text" id="input_number">
        <input type="submit" onClick="getTwice(); return false;" value="送信">
    </form>

    <div id="result">
    </div>
</body>
</html>

<script>
function getTwice() {
    const inputNumber = document.getElementById('input_number').value;
    const url = '/twice/' + inputNumber;
    fetch(url).then(function(response) {
        return response.json();
    }).then(function(json) {
        var result = document.querySelector('#result');
        if (json.message != "") {
            result.innerHTML = json.message;
            return;
        }

        result.innerHTML = json.output;
    });
}
</script>
{{end}}
