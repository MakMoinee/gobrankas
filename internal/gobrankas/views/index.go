package views

var HomeView = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body background="color: #d5d5d5;">
    <form action="/upload" method="post" enctype="multipart/form-data">
        <input type="file" name="file" id="mfile">
        <br>
        <br>
        <button type="submit">Upload file</button>
    </form>
</body>
</html>`
