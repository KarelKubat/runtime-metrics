test:
	for d in `find * -type d | grep -v tools`; do \
	  make -C $$d test; \
	done
