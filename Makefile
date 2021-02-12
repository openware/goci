APPS     :=	goci

.PHONY: $(APPS)

all:	$(APPS)

$(APPS):
	CGO_ENABLED=0 go build --ldflags "-X main.Version=$$(cat .tags)" -o bin/$@ ./cmd/$@

clean:
	rm -rf bin/*
