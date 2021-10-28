#!/bin/bash

# PATH FOR THE EXECUTABLES
BUILDPATH="./bin"

# GET EITHER TAG OR COMMIT HASH (SHORT)
VERSION=$(git describe --tags)
if [ "$VERSION" == "" ]; then
    COMMIT=$(git rev-parse HEAD)
    VERSION="dev_${COMMIT:0:6}"
fi

# CREATE FOLDER FOR EXECUTABLES (IF IT DOESN'T EXIST)
if [ ! -d $BUILDPATH ]; then
    mkdir $BUILDPATH
fi

PLATFORMS=(
    'linux;arm'
    'linux;amd64'
    'windows;amd64'
    'darwin;amd64'
)

# BUILD FRONT END FIRST SINCE IT CAN BE THE SAME FOR ALL PLATFORMS
echo "🔨 Building front end..."
cd ../web
npm install
npm run build
cd ../scripts

# MOVE FRONT END ASSETS TO BUILDPATH
cd $BUILDPATH
mkdir ./web/
mv ../../web/build ./web/
cd ../


# BUILD BACK END FOR ALL PLATFORMS AND CREATE ARCHIVE WITH THE EXECUTABLE
# AND THE FRONT END ASSETS
for PLATFORM in ${PLATFORMS[*]}; do
    IFS=';' read -ra SPLIT <<< "$PLATFORM"
    OS=${SPLIT[0]}
    ARCH=${SPLIT[1]}
    EXECUTABLE=goshortrr_${VERSION}_${OS}_${ARCH}
    EXECUTABLE_PATH=${BUILDPATH}/${EXECUTABLE}

    echo "🔨 Building ${OS}_$ARCH..."
    (env GOOS=$OS GOARCH=$ARCH \
        go build -v -o ${EXECUTABLE_PATH} \
        ../cmd/goshortrr)

    cd $BUILDPATH
    # COMPRESS EXECUTABLE AND FRONT END ASSETS INTO ONE ARCHIVE
    tar -cf ./${EXECUTABLE}.tar.gz ./${EXECUTABLE} ./web/build/
    # GENERATE CHECKSUM FOR ARCHIVE
    sha256sum ./${EXECUTABLE}.tar.gz | tee ./${EXECUTABLE}_sha256sum.txt
    # md5sum ./${EXECUTABLE}.tar.gz | tee ./${EXECUTABLE}_md5sum.txt
    rm ${EXECUTABLE}
    cd ..

done

# DELETE FRONT END ASSETS FROM BUILDPATH
cd $BUILDPATH
rm -r ./web
cd ..

echo ""
echo "BUILDS DONE! 🥳"

wait