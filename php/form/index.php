<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>PHP form</title>
</head>

<body>
  <form method="POST">
    <input type="text" id="name" name="name" placeholder="Enter your name" required>
    <input type="email" id="email" name="email" placeholder="Enter your email" required>
    <input type="password" id="password" name="password" placeholder="Enter your password" required>
    <button type="submit" value="submit" id="submit-btn">Enter</button>
  </form>
  <?php
  $name = htmlspecialchars($_POST["name"]);
  $email = filter_var($_POST["email"], FILTER_SANITIZE_EMAIL);
  $password = $_POST["password"];
  $hashedPassword = password_hash($password, PASSWORD_ARGON2ID);
  $newUser = [$name, $email, $hashedPassword];
  if ($name == "" || $password == "" || $email == "") {
    echo "Error, one of the spaces is empty";
  } else {
    $usersContent = [json_decode(file_get_contents("users.json"))];
    if ($usersContent == "") {
      $usersArray = [0 => $newUser];
      file_put_contents("users.json", json_encode($usersArray));
    } else {
      $usersContent[-1] = $newUser;
      file_put_contents("users.json", json_encode($usersContent));
    }
  }

  ?>
</body>

</html>