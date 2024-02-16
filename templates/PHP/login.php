<?php
    if(session_status() == PHP_SESSION_NONE) {
        session_start();
    }
    require("path/to/mysql.php");

    if(isset($_POST["login_button"])) {
        $stmt = $con->prepare("SELECT * FROM users WHERE email=:email");
        $stmt->bindParam(":email", $_POST["login_email"]);
        $stmt->execute();
        $userexist = $stmt->fetchAll();

        $passwordhash = $userexist[0]["password"];
        $chkpw = password_verify($_POST["login_password"], $passwordhash);

        if($chkpw == true) {
            $_SESSION["ID"] = $userexist[0]["id"];
            $_SESSION["firstname"] = $userexist[0]["firstname"];
            // header("Location: ../index.php");
        }else {
            echo "Password invalid";
        }
    }

?>
<style>

    html,body {
        height: 100%;
    }

    body {
        display: flex;
        align-items: center;
        padding-top: 40px;
        padding-bottom: 40px;
    }

    .form-signin {
        max-width: 330px;
        padding: 15px;
    }

    .form-signin .form-floating:focus-within {
        z-index: 2;
    }

    a {
        text-decoration: none;
    }
</style>
<main class="form-signin w-100 m-auto">
    <form method="post">
        <h1 class="h3 mb-3 fw-normal">Sign In</h1>

        <div class="form-floating" style="color: black;">
            <input type="email" class="form-control" id="login_email" name="login_email">
            <label for="login_email">Email</label>
        </div>
        <br>
        <div class="form-floating" style="color: black;">
            <input type="password" class="form-control" id="login_password" name="login_password">
            <label for="login_password">Password</label>
        </div>
        <hr>
        <button class="w-100 btn custommetelbutton" name="login_button" type="submit">Sign In</button>
        <div>
            <p class="text-center">&copy;2024 ProjectName <a href="register.php">Need an Account?</a></p>
        </div>
    </form>
</main>  