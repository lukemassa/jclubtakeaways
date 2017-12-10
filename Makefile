html:
	cd src/templates; j2 index.html > $(CURDIR)/docs/index.html
	cd src/templates; j2 weeks.html > $(CURDIR)/docs/weeks.html
	cd src/templates; j2 about.html > $(CURDIR)/docs/about.html
	cd src/templates; j2 oneweek.html > $(CURDIR)/docs/oneweek.html
	cd src/templates; j2 upcomingweeks.html > $(CURDIR)/docs/upcomingweeks.html
