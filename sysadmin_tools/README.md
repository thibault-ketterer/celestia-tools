# celestia admin tools for nodes
purpose: auto upgrade celestia node (work for light, maybe other types)

- install the systemd file (/!\ overwrite it)
- stop previous celestia node
- checkout last version of celestia (based on tags order by version, sort -V)
- compiles it
- installs it
- upgrade config if needed (light only)
- start the service again
- show status
- shows version

It looks like that


    $ bash upgrade_celestia.sh
    # updating system config...
    # stopping daemon...
    
    last tag is [v0.9.5]
    
    # checkout...
    
    # building from source if needed...
    --> Building Celestia
    # installing...
    --> Installing Celestia
    removed '/usr/local/bin/celestia'
    './build/celestia' -> '/usr/local/bin/celestia'
    
    # update config...
    
    # start service...
    
    ● celestia-lightd.service - celestia-lightd Light Node
         Loaded: loaded (/etc/systemd/system/celestia-lightd.service; enabled; vendor preset: enabled)
         Active: active (running) since Mon 2023-05-15 18:47:32 UTC; 6ms ago
       Main PID: 3055933 (start.sh)
          Tasks: 4 (limit: 38251)
         Memory: 2.5M
            CPU: 4ms
         CGroup: /system.slice/celestia-lightd.service
                 ├─3055933 /bin/bash /home/ubuntu/celestia-node/start.sh
                 ├─3055936 timeout 4 curl -s https://rpc-blockspacerace.pops.one
                 └─3055937 curl -s https://rpc-blockspacerace.pops.one
    
    May 15 18:47:32 node5 systemd[1]: Started celestia-lightd Light Node.
    May 15 18:47:32 node5 start.sh[3055933]: start
    May 15 18:47:32 node5 start.sh[3055934]:     __                    _     _            _
    May 15 18:47:32 node5 start.sh[3055934]:    / / __ _ __   ___     | |__ | | ___   ___| | _____ _ __   __ _  ___ ___
    May 15 18:47:32 node5 start.sh[3055934]:   / / '__| '_ \ / __|____| '_ \| |/ _ \ / __| |/ / __| '_ \ / _` |/ __/ _ \
    May 15 18:47:32 node5 start.sh[3055934]:  / /| |  | |_) | (_|_____| |_) | | (_) | (__|   <\__ \ |_) | (_| | (_|  __/
    May 15 18:47:32 node5 start.sh[3055934]: /_/ |_|  | .__/ \___|    |_.__/|_|\___/ \___|_|\_\___/ .__/ \__,_|\___\___|
    May 15 18:47:32 node5 start.sh[3055934]:          |_|                                         |_|
    
    Semantic version: v0.9.5
    Commit: 2fa72c7199e5b93772a2c7e25141cfbd28f16a8e
    Build Date: Mon May 15 18:47:28 UTC 2023
    System version: amd64/linux
    Golang version: go1.20.3
    
    upgrade finished
    have a good day


# start and loop over all known nodes
purpose: loop over all nodes in case of full node failure

    $ bash start.sh

replace service in systemd with call to start.sh, see above

# small utils

## light node data cleaner
purpose: save from miscommand / mistakes on light nodes
    $ bash clean_data.sh


