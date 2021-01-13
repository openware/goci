APPS     :=	goci

.PHONY: $(APPS)

all:	$(APPS)

$(APPS):
	go build --ldflags "-X main.Version=$$(cat .tags)" -o bin/$@ ./cmd/$@

clean:
	rm -rf bin/*
