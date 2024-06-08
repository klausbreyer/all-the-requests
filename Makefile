TAG_PREFIX = v

# Extract the latest tag from Git
CURRENT_VERSION = $(shell git describe --tags `git rev-list --tags --max-count=1`)
CURRENT_MAJOR = $(shell echo $(CURRENT_VERSION) | cut -d. -f1 | sed 's/$(TAG_PREFIX)//')
CURRENT_MINOR = $(shell echo $(CURRENT_VERSION) | cut -d. -f2)
CURRENT_PATCH = $(shell echo $(CURRENT_VERSION) | cut -d. -f3)

.PHONY: tag_major tag_minor tag_patch

# Update the major version
tag_major:
	@NEW_MAJOR=$$(($(CURRENT_MAJOR) + 1)); \
	echo "New Major Version: $$NEW_MAJOR.0.0"; \
	NEW_VERSION=$$NEW_MAJOR.0.0; \
	git tag -a $(TAG_PREFIX)$$NEW_VERSION -m "Release $(TAG_PREFIX)$$NEW_VERSION"; \
	git push origin $(TAG_PREFIX)$$NEW_VERSION

# Update the minor version
tag_minor:
	@NEW_MINOR=$$(($(CURRENT_MINOR) + 1)); \
	echo "New Minor Version: $(CURRENT_MAJOR).$$NEW_MINOR.0"; \
	NEW_VERSION=$(CURRENT_MAJOR).$$NEW_MINOR.0; \
	git tag -a $(TAG_PREFIX)$$NEW_VERSION -m "Release $(TAG_PREFIX)$$NEW_VERSION"; \
	git push origin $(TAG_PREFIX)$$NEW_VERSION

# Update the patch version
tag_patch:
	@NEW_PATCH=$$(($(CURRENT_PATCH) + 1)); \
	echo "New Patch Version: $(CURRENT_MAJOR).$(CURRENT_MINOR).$$NEW_PATCH"; \
	NEW_VERSION=$(CURRENT_MAJOR).$(CURRENT_MINOR).$$NEW_PATCH; \
	git tag -a $(TAG_PREFIX)$$NEW_VERSION -m "Release $(TAG_PREFIX)$$NEW_VERSION"; \
	git push origin $(TAG_PREFIX)$$NEW_VERSION
