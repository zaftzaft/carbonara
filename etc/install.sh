service=carbonara

systemctl is-enabled $service
if [ $? -eq 0 ]; then
  systemctl stop $service
fi

if [ -d /usr/lib/systemd/system/ ]; then
  unit_dir=/usr/lib/systemd/system
else
  unit_dir=/etc/systemd/system
fi


cd ../
go build
cd -

install -Dm 644 "./carbonara.service" "${unit_dir}/${service}.service"
install -Dm 755 "../carbonara" "/usr/bin/carbonara"



systemctl daemon-reload
systemctl enable $service
systemctl start $service

