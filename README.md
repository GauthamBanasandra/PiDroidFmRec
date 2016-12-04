# PiDroidFmRec
PiDroidFmRec is a **Go** framework for recording FM on **Android** devices.<br/>
### How does it work?
1) Android exposes a command-line tool **"[adb](https://developer.android.com/studio/command-line/adb.html)"**. <br/>
2) Adb is used to send input events to the phone such as tapping on an (x, y) location, sending text input, pressing a key and so on. <br/>
3) Thus, there's a configuration file that must be filled out by the user and the Go framework will provide hooks as to when the actions will be executed. <br/>
4) The Go code will run on a separate computer (**Raspberry PI**)
### Why not just Android?
Adb by itself can't run on the Android device. A workaround would be to root the Android device but this isn't viable as it requires some expertise in doing so and also it may void the warranty of the device.<br/>
### Why use adb?
1) Android API for FM radio is not available. It's upto the manufacturers to provide an API for the FM module. <br/>
2) Even if the API exists, it will be different for each device model. <br/>
3) Hence, it's not a scalable solution to just use the API because of the non-uniformity in the API. <br/>
4) However, adb is accessible across all the Android device models, which is supported even in the older versions of Android. <br/>
##
### Installation
**Installing adb**<br/>
>Linux
>```sh
>$ sudo apt-get install android-tools-adb
>```
>[Click here](http://www.howtogeek.com/125769/how-to-install-and-use-abd-the-android-debug-bridge-utility/) for Windows.<br/>

**Running adb over Wi-Fi**
The host for running the Go code (computer) is connected to the Android device over Wi-fi. To turn this functionality on, follow these steps - <br/>
>First connect the Android device to the computer with a USB cable and execute -
>```sh
>$ adb tcpip 5555
>```
>Unplug the device and then -
>```sh
>$ adb connect <device_ip>:5555
>```
>The above procedure must be done everytime the Android device restarted.


### Using the framework
As mentioned earlier, all that the user must do is fill out the configuration file and the Go framework will translate the configuration into adb commands and execute them accordingly. [Click here](https://github.com/GauthamBanasandra/PiDroidFmRec/blob/master/Go/config.json) for configuration sample.<br/>

The configuration comprises of 3 main parts - <br/>
1) **deviceInfo** (Device information) - <br/>
Contains the configuration regarding the **IP** address of the device and the **device password** needed to unlock the phone.
```
"deviceInfo": {
        "ip": "192.168.0.2",
        "devicePwd": "123456"
    }
```
2) **action** - <br/>
A list of actions that needs to be executed on the Android device. Each action consists of a **cmd** (command) and its **input**.
```
"action": [
        {
            "cmd": "tap",
            "input": {
                "x1": 535,
                "y1": 1695
            }
        },
        {
            "cmd": "text",
            "input": {
                "text": "Devotional hour 91.9"
            }
        },
        {
            "cmd": "keyevent",
            "input": {
                "key": 66
            }
        }
    ]
```
>**input** can take the following values - <br/>
>- x1, y1 for **tap** command.
>- x1, y1, x2, y2 for **swipe** command.
>- text for **text** command.
>- packageName for **monkey** command.<br/>

[Click here](https://developer.android.com/studio/command-line/adb.html#issuingcommands) for the complete list of adb commands (Please note that all the adb commands are not supported by the framework currently).<br/>
To show the co-ordinates of the tap on the Android device, go to -
>Settings > Developer options > Pointer location

3) **recordInfo** (Record information) - <br/>
Contains information for recording FM. <br/>
Recording will be scheduled to begin **startTime** and end at **stopTime**. The times must be specified in 24 hour clock format only. Alarms are scheduled to trigger at both these times and the corresponding actions specified by **startActionIdx** and **stopActionIdx** will be executed.<br/>
The actions listed under **action** in the configuration will be executed according to the order specified by the **startActionIdx** and **stopActionIdx**. A Zero-based indexing is followed.
```
"recordInfo": {
        "startTime": {
            "hh": 21,
            "mm": 23,
            "ss": 50
        },
        "stopTime": {
            "hh": 21,
            "mm": 28,
            "ss": 0
        },
        "startActionIdx": [
            0,
            1
        ],
        "stopActionIdx": [
            1,
            2,
            3,
            3
        ]
    }
```
### External components
This project includes a companion Android app.<br/>
**Features**
1) Depending on the Android device, unlocking it via adb takes nearly a minute to execute. The Android app in this project has a functionality to schedule the device to wake up at a particular time. <br/>
2) The app lists the name and the corresponding package names of all the apps installed on the device, which could be used in writing the configuration. <br/>
3) The app displays the IP address of the device. <br/>
### Extensibility
This project isn't just restricted to recording FM on Android devices. It could help one automate any task with the appropriate configuration.