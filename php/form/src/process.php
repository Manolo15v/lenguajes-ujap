<?php
// Verificar que se haya enviado el formulario por POST
if ($_SERVER["REQUEST_METHOD"] != "POST") {
    header("Location: index.php");
    exit();
}

// 1. VALIDACIÓN: Verificar que todos los campos no estén vacíos
if (empty($_POST['nombre']) || empty($_POST['email']) || empty($_POST['password'])) {
    $error_message = "Error: Todos los campos son obligatorios.";
    header("Location: index.php?error=" . urlencode($error_message));
    exit();
}

// Obtener y limpiar los datos del formulario
$nombre = trim($_POST['nombre']);
$email = trim($_POST['email']);
$password_plana = $_POST['password'];

// 2. HASHING DE CONTRASEÑA: Utilizar password_hash()
$password_hash = password_hash($password_plana, PASSWORD_DEFAULT);

// 3. PREPARACIÓN DE DATOS: Crear array asociativo con los datos del nuevo usuario
$nuevo_usuario = [
    'nombre' => $nombre,
    'email' => $email,
    'password_hash' => $password_hash
];

// 4. ALMACENAMIENTO EN JSON
$archivo = 'usuarios.json';

// Leer los datos existentes
$usuarios = [];
if (file_exists($archivo)) {
    $contenido_json = file_get_contents($archivo);
    if (!empty($contenido_json)) {
        $usuarios = json_decode($contenido_json, true);
        if (!is_array($usuarios)) {
            $usuarios = [];
        }
    }
}

// Agregar el nuevo usuario al array
$usuarios[] = $nuevo_usuario;

// Guardar los datos en el archivo JSON
$json_para_guardar = json_encode($usuarios, JSON_PRETTY_PRINT);

if (file_put_contents($archivo, $json_para_guardar)) {
    $success_message = "¡Usuario registrado con éxito!";
    header("Location: index.php?success=" . urlencode($success_message));
} else {
    $error_message = "Error: No se pudo escribir en el archivo JSON. Verifique los permisos.";
    header("Location: index.php?error=" . urlencode($error_message));
}
exit();
?>