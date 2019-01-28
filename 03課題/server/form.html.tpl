<!DOCTYPE html>
<html>
<body>
    <form action="./" method="get">
        <input type="text" name="input_number">
        <input type="submit" onClick="getTwice(); return false;" value="送信">
    </form>

    <div id="result">
    </div>
</body>
</html>

<script>
function getTwice() {
    const url = 'http://localhost:9090/twice/2';
    fetch(url).then(function(response) {
        return response;
    }).then(function(json) {
        var result = document.querySelector('#result');
        result.innerHTML = json.output;
    });
}
</script>