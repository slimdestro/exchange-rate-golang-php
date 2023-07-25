<?php
    /***
	- api.php: serves frontend/index.php 
    - curl is required otherwise the backend(golang) will need some extra header to make it open for all origins(CORS error)
	- Author: Mukul(https://github.com/slimdestro) | https://www.modcode.dev
    */
    $api_url = 'http://localhost:8080/frontend';
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, $api_url);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
    $response = curl_exec($ch);
    curl_close($ch);

    if (!$response) {
        header("HTTP/1.1 500 Internal Server Error");
        die('Error fetching data from the API.');
    }

    header('Content-Type: application/json');
    echo $response;

?>