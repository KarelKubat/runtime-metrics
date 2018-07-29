test:
	for d in `find * -type d | grep -v tools`; do \
	  echo "------- Running tests in $$d -------"; \
	  make -C $$d test || exit 1; \
	done
