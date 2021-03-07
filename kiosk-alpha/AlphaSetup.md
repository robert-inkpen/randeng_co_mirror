## Alpha Setup

This setup uses **Raspberry Pi OS Lite**,  **Xserver** for display and **Chromium** as the browser. The below steps accomplish a rough alpha that will deploy a local server and displays the page in kiosk mode.

**TODO:** Manage the setup through a scripting task

### Steps 

- `raspi-config`
  - timezone
  - user pwd
  - Desktop/CLI -> Console Autologin
  - Enable SSH
  - Disable Overscan
- enable wifi  (/etc/wpa_supplicant/wpa_supplicant/conf)
  - overwrite above file after reading wifi details from user
- `sudo apt-get update`
- `sudo apt-get upgrade`
- `sudo apt-get install --no-install-recommends xserver-xorg x11-xserver-utils xinit openbox`
- `sudo apt-get install --no-install-recommends chromium-browser`

**Xserver Openbox Config** (/etc/xdg/openbox/autostart )

```bash
xset s off
xset s noblank
xset -dpms

chromium-browser --disable-infobars --kiosk 'http://localhost:3000'
```

**NodeJs** : `wget -O - https://raw.githubusercontent.com/meech-ward/NodeJs-Raspberry-Pi/master/Install-Node.sh | sudo bash`

**Deploy** : bash script to launch server and run xserver display

- go/node server launch
- `startx -- -nocursor`