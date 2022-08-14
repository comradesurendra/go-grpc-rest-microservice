# GRPC - Template

**To install latest version of Protoc**

    sudo apt-get remove protobuf-compiler

    PB_REL="https://github.com/protocolbuffers/protobuf/releases"

    curl -LO $PB_REL/download/v3.19.1/protoc-3.19.1-linux-x86_64.zip

**For linux**

    sudo apt install unzip

    PROTOC_ZIP=protoc-3.19.1-linux-x86_64.zip

    sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc

    sudo unzip -o $PROTOC_ZIP -d /usr/local include/*

    chmod +x /usr/local/bin/protoc

    rm -f $PROTOC_ZIP

**We have used go mod name GRPC-SERVE (You can replace mod name by replacing in source code usually called most in import)**

For updating go mod please check go documentation. Basic command here to start from new

    go mod init GRPC-SERVE

__Go install require package mentioned in tools.go Details instruction at : https://github.com/grpc-ecosystem/grpc-gateway#installation__

    go mod tidy

    go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest \
    google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

**Generate compiled go code for grpc gateway**

    make file_name=home
    here (home) is the proto name

**HTTP and RPC server configuration**

    Update config/config.yaml file

**Start gRPC Server**

    go run server/main.go

**start reverse proxe server**

    go run reverse-proxy/main.go


**localhost url**

    ENDPOINT: http://localhost:port/v1/loan/homedata?value=test

  
**Curl check**

    curl -d "value=test data again" http://localhost:port/v1/loan/homedata?value=test

  

## Run SwaggerUI :-

    go to the swaggerui dir and run cmd
    go run main.go
    once server starts go to the browser and hit 
    http://localhost:8080

## GoDoc :-

**GoDoc Creation**

Example :: creating go doc for _``file.go``_

    // Package <PackageName> <purpose of package>

    package <PackageName>

    // Function <FunctionName> <Function type>
    // Request *PackageRequest as request and 
    // Response *PackageResponse, error

    func (s *PackageStruct) ReceiverFunctionName(ctx context.Context, req *PackageRequest) (*PackageResponse, error) {

        return &PackageResponse, nil
    }

running godoc

    godoc -http=:6060

**make sure port 6060 is free or use the suitable port**

open browser

    http://localhost:6060/pkg/GRPC-SERVE

**!! you will get all your docs here !!**


## Zap (Level logging) :-

**For Development Environment**

    logger, _ := zap.NewDevelopment()
    logger.Debug("debug")

Debug only works in development environment

**For Production Environment**

    logger, _ := zap.NewProduction() -->

**Level**

    helper.SugarObj.Info("Info Level")
    helper.SugarObj.Error("Error")
    helper.SugarObj.Warn("Warning")


## Testing :-

• To write unit tests in GoLang, we need to import the testing package. There are some rules when doing testing. The filename must end with _test.

        http_test.go

• The command to be used is ``go test`` in the CLI. It will test all the files set up as test files and return the output in the terminal.

• The functions in that test file must follow this signature.

        Func TestName(t *Testing.T){
        //code for testing
        }

• Try running-

       go test -cover

_You should see-_

    PASS
    coverage: 100.0% of statements

**BENCHMARKS:-**

• Function 

        func BenchmarkXxx(*testing.B){
           //code for testing/benchmark.
         }

• Benchmarks are executed by the ``go test`` command when its -bench flag is provided. Benchmarks are run sequentially.

• A sample Benchmark function look like this-

        func BenchmarkRandInt(b *testing.B) {
            for i := 0; i < b.N; i++ {
                rand.Int()
            }

• If Benchmark need some expensive settings , the timer must be reset;-

        b.ResetTimer()  

• When running the go test we must provide the -bench flag which will show the output.

        // the dot is the regex matching everything
        go test -bench=.    
    
• If a benchmark needs to test performance in a parallel setting, it may use the ``RunParallel helper function;`` such benchmarks are intended to be used with the 

        go test -cpu flag:

**SKIPPING SOME TEST/BENCHMARKS:-**

• Tests or benchmarks may be skipped at run time with a call to the ``Skip method of *T or *B``:

• eg:-

    func TestTimeConsuming(t *testing.T) {
        if testing.Short() {
            t.Skip("skipping test in short mode.")
        }
        ...
    }
