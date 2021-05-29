<html>
    <head>
        <title>Ingest</title>
    </head>
    <body>
        <?php

            $json = file_get_contents('php://input');

            // decode the json data
            $data = json_decode($json);

            // will be "null" if json is not valid
            $stringified_data = json_encode($data);

            $redis = new Redis();
            //Connecting to Redis
            $redis->connect('localhost', 6379);
            
            if ($stringified_data != "null"){
                echo "Valid";
                $redis->lPush("postback_queue", $stringified_data);
            }
            else {
                echo "Invalid";
            }
    
            
            
        
        ?>
    </body>
</html>