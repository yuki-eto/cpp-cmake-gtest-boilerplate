FROM multiarch/crossbuild

RUN apt-get remove -y cmake
RUN wget https://github.com/Kitware/CMake/releases/download/v3.17.0/cmake-3.17.0-Linux-x86_64.sh
RUN sh cmake-3.17.0-Linux-x86_64.sh --prefix=/usr/local --skip-license
RUN rm cmake-3.17.0-Linux-x86_64.sh
