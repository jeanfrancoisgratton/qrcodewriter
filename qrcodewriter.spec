%define debug_package   %{nil}
%define _build_id_links none
%define _name   qrcodewriter
%define _prefix /opt
%define _version 0.200
%define _rel 0
%define _arch x86_64
%define _binaryname qrcw

Name:       qrcodewriter
Version:    %{_version}
Release:    %{_rel}
Summary:    qrcodewriter

Group:      SSL
License:    GPL3.0
URL:        https://github.com/jeanfrancoisgratton/qrcodewriter

Source0:    %{name}-%{_version}.tar.gz
BuildArchitectures: x86_64
BuildRequires: gcc
#Requires: sudo
#Obsoletes: vmman1 > 1.140

%description
Creates a QR code from a user-supplied URL

%prep
%autosetup

%build
cd %{_sourcedir}/%{_name}-%{_version}/src
PATH=$PATH:/opt/go/bin go build -o %{_sourcedir}/%{_binaryname} .
strip %{_sourcedir}/%{_binaryname}

%clean
rm -rf $RPM_BUILD_ROOT

%pre
exit 0

%install
install -Dpm 0755 %{_sourcedir}/%{_binaryname} %{buildroot}%{_bindir}/%{_binaryname}

%post

%preun

%postun

%files
%defattr(-,root,root,-)
%{_bindir}/%{_binaryname}


%changelog
* Wed Jun 21 2023 builder <builder@famillegratton.net> 0.200-0
- Subcommand and output fixes (jean-francois@famillegratton.net)
- Doc fix (builder@famillegratton.net)

* Wed Jun 21 2023 builder <builder@famillegratton.net> 0.100-0
- new package built with tito

