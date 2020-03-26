FROM multiarch/crossbuild

RUN apt-get remove -y cmake
RUN wget https://github.com/Kitware/CMake/releases/download/v3.17.0/cmake-3.17.0-Linux-x86_64.sh
RUN sh cmake-3.17.0-Linux-x86_64.sh --prefix=/usr/local --skip-license
RUN rm cmake-3.17.0-Linux-x86_64.sh

RUN apt-get install -y valgrind libpcre3-dev

RUN wget http://prdownloads.sourceforge.net/swig/swig-4.0.1.tar.gz
RUN tar xvf swig-4.0.1.tar.gz
RUN cd swig-4.0.1 && ./configure && make && make install
RUN rm -rf ./swig-4.0.1*
