all: test docs

test:
	@for d in `find * -type d | grep -v tools`; do \
	  echo "------- Running tests in $$d -------"; \
	  make -C $$d test || exit 1; \
	done

docs:
	@for d in . `find * -type d | grep -v tools`; do \
	  echo "------- Making docs in $$d -------"; \
	  godocdown $$d > $$d/README.md || exit 1; \
	done
