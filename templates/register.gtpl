<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Register</title>
    <style>
        body{
            /*background: url("./images/background_main.jpg"); /* Путь к фоновому изображению */*/
        }
     </style>
</head>
<body>

<form method="POST" action="">
    Username: <input type="text" name="username"  /><br />
    Password: <input type="password" name="password"  /><br />
    Email: <input type="email" name="email"  /><br />
    First Name: <input type="text" name="first_name"  /><br />
    Last Name: <input type="text" name="last_name"  /><br />
    <div align="center">
        <p><input type="submit" value="Register" /></p>
    </div>
</form>

{{ . }}

{{template "bar" .}}


</br>

</body>


</html>