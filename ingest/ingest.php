<html>
    <head>
        <title>Ingest</title>
    </head>
    <body>
        <?php
    
            $redis = new Redis();
            //Connecting to Redis
            $redis->connect('localhost', 6379);
            // $redis->auth('password');
        
            if ($redis->ping()) {
                echo "PONGn";
            }
        
        ?>
    </body>
</html>