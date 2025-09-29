#!/bin/sh

if [ -e "$SOCKET_PATH" ]; then
	rm -f "$SOCKET_PATH"
fi

chown -R formpf:formpf /var/log/form-pf $STORAGE_PATH 
exec su formpf -c "$@"

