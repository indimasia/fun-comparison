<?php
$host = 'localhost:33061';
$db   = 'comparison';
$user = 'root';
$pass = '';
$charset = 'utf8mb4';

$dsn = "mysql:host=$host;dbname=$db;charset=$charset";
$options = [
    PDO::ATTR_ERRMODE            => PDO::ERRMODE_EXCEPTION,
    PDO::ATTR_DEFAULT_FETCH_MODE => PDO::FETCH_ASSOC,
    PDO::ATTR_EMULATE_PREPARES   => false,
];

$pdo = new PDO($dsn, $user, $pass, $options);
try {
    $pdo->beginTransaction();
    $startTime = microtime(true);

    $stmt = $pdo->prepare('INSERT INTO employees (name, salary, greeting) VALUES (?, ?, ?)');

    for ($i = 0; $i < 100000; $i++) {
        $name = substr(str_shuffle(str_repeat($x='0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ', ceil(10/strlen($x)) )),1,10);
        $salary = rand(10000, 100000);
        $greeting = (rand(0, 1) == 0) ? 'Mr' : 'Ms';
        $stmt->execute([$name, $salary, $greeting]);
    }

    $stmt = $pdo->prepare('INSERT INTO employees2 (name, salary, greeting) VALUES (?, ?, ?)');

    for ($i = 0; $i < 100000; $i++) {
        $name = substr(str_shuffle(str_repeat($x='0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ', ceil(10/strlen($x)) )),1,10);
        $salary = rand(10000, 100000);
        $greeting = (rand(0, 1) == 0) ? 'Mr' : 'Ms';
        $stmt->execute([$name, $salary, $greeting]);
    }

    $pdo->commit();
    $endTime = microtime(true);
    echo "PHP script execution time: " . ($endTime - $startTime) . " seconds\n";
} catch (\PDOException $e) {
    $pdo->rollBack();
    throw new \PDOException($e->getMessage(), (int)$e->getCode());
}
?>
