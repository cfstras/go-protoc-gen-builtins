FROM ghcr.io/wasilibs/wasix-sdk:sha-fc94d60
#FROM --platform=linux/amd64 ghcr.io/wasilibs/wasix-sdk:sha-fc94d60

RUN apt-get update && apt-get install -y binaryen git patch

ADD buildtools/wasm/version.txt version.txt
RUN git clone --recursive https://github.com/protocolbuffers/protobuf.git --branch $(cat version.txt | awk '{$1=$1};1') /workspace

ENV CXXFLAGS "${CXXFLAGS} -O3 -pthread -D_WASI_EMULATED_MMAN -D_WASI_EMULATED_PROCESS_CLOCKS -DABSL_HAVE_MMAP"
ENV LDFLAGS $LDFLAGS -lwasi-emulated-mman -lwasi-emulated-process-clocks -Wl,--max-memory=4294967296

WORKDIR /workspace

# From https://github.com/protocolbuffers/protobuf/blob/v25.1/src/google/protobuf/compiler/cpp/main.cc

ADD buildtools/wasm/main.cc main.cc
ADD buildtools/wasm/create_plugin_entrypoints.py create_plugin_entrypoints.py
RUN mkdir -p cmake/src && python3 create_plugin_entrypoints.py

WORKDIR /workspace/cmake

RUN cmake . -B build -Dprotobuf_BUILD_TESTS=OFF
RUN cmake --build build --target plugins --target protoc -j$(nproc)

SHELL ["/bin/bash", "-c"]
RUN for f in build/protoc-gen-*[^0-9] build/protoc; do wasm-opt -o "$f".wasm --low-memory-unused --flatten --rereloop --converge -O3 "$f"; done

CMD ["bash", "-c", "cp build/*.wasm /out/"]
