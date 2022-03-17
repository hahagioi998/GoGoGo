// Code generated by thriftgo (0.1.7). DO NOT EDIT.

package snow

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

type SnowRequest struct {
	Wanted int64 `thrift:"wanted,1" json:"wanted"`
}

func NewSnowRequest() *SnowRequest {
	return &SnowRequest{}
}

func (p *SnowRequest) GetWanted() (v int64) {
	return p.Wanted
}
func (p *SnowRequest) SetWanted(val int64) {
	p.Wanted = val
}

var fieldIDToName_SnowRequest = map[int16]string{
	1: "wanted",
}

func (p *SnowRequest) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_SnowRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *SnowRequest) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.Wanted = v
	}
	return nil
}

func (p *SnowRequest) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("SnowRequest"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *SnowRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("wanted", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.Wanted); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *SnowRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SnowRequest(%+v)", *p)
}

func (p *SnowRequest) DeepEqual(ano *SnowRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Wanted) {
		return false
	}
	return true
}

func (p *SnowRequest) Field1DeepEqual(src int64) bool {

	if p.Wanted != src {
		return false
	}
	return true
}

type SnowResponse struct {
	SnowCount int64 `thrift:"snowCount,1" json:"snowCount"`
}

func NewSnowResponse() *SnowResponse {
	return &SnowResponse{}
}

func (p *SnowResponse) GetSnowCount() (v int64) {
	return p.SnowCount
}
func (p *SnowResponse) SetSnowCount(val int64) {
	p.SnowCount = val
}

var fieldIDToName_SnowResponse = map[int16]string{
	1: "snowCount",
}

func (p *SnowResponse) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_SnowResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *SnowResponse) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.SnowCount = v
	}
	return nil
}

func (p *SnowResponse) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("SnowResponse"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *SnowResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("snowCount", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.SnowCount); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *SnowResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SnowResponse(%+v)", *p)
}

func (p *SnowResponse) DeepEqual(ano *SnowResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.SnowCount) {
		return false
	}
	return true
}

func (p *SnowResponse) Field1DeepEqual(src int64) bool {

	if p.SnowCount != src {
		return false
	}
	return true
}

type Snow interface {
	Wanted(ctx context.Context, req *SnowRequest) (r *SnowResponse, err error)
}

type SnowClient struct {
	c thrift.TClient
}

func NewSnowClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *SnowClient {
	return &SnowClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewSnowClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *SnowClient {
	return &SnowClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewSnowClient(c thrift.TClient) *SnowClient {
	return &SnowClient{
		c: c,
	}
}

func (p *SnowClient) Client_() thrift.TClient {
	return p.c
}

func (p *SnowClient) Wanted(ctx context.Context, req *SnowRequest) (r *SnowResponse, err error) {
	var _args SnowWantedArgs
	_args.Req = req
	var _result SnowWantedResult
	if err = p.Client_().Call(ctx, "wanted", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type SnowProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      Snow
}

func (p *SnowProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *SnowProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *SnowProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewSnowProcessor(handler Snow) *SnowProcessor {
	self := &SnowProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("wanted", &snowProcessorWanted{handler: handler})
	return self
}
func (p *SnowProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type snowProcessorWanted struct {
	handler Snow
}

func (p *snowProcessorWanted) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := SnowWantedArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("wanted", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := SnowWantedResult{}
	var retval *SnowResponse
	if retval, err2 = p.handler.Wanted(ctx, args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing wanted: "+err2.Error())
		oprot.WriteMessageBegin("wanted", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("wanted", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type SnowWantedArgs struct {
	Req *SnowRequest `thrift:"req,1" json:"req"`
}

func NewSnowWantedArgs() *SnowWantedArgs {
	return &SnowWantedArgs{}
}

var SnowWantedArgs_Req_DEFAULT *SnowRequest

func (p *SnowWantedArgs) GetReq() (v *SnowRequest) {
	if !p.IsSetReq() {
		return SnowWantedArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *SnowWantedArgs) SetReq(val *SnowRequest) {
	p.Req = val
}

var fieldIDToName_SnowWantedArgs = map[int16]string{
	1: "req",
}

func (p *SnowWantedArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SnowWantedArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_SnowWantedArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *SnowWantedArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Req = NewSnowRequest()
	if err := p.Req.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *SnowWantedArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("wanted_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *SnowWantedArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Req.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *SnowWantedArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SnowWantedArgs(%+v)", *p)
}

func (p *SnowWantedArgs) DeepEqual(ano *SnowWantedArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Req) {
		return false
	}
	return true
}

func (p *SnowWantedArgs) Field1DeepEqual(src *SnowRequest) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

type SnowWantedResult struct {
	Success *SnowResponse `thrift:"success,0" json:"success,omitempty"`
}

func NewSnowWantedResult() *SnowWantedResult {
	return &SnowWantedResult{}
}

var SnowWantedResult_Success_DEFAULT *SnowResponse

func (p *SnowWantedResult) GetSuccess() (v *SnowResponse) {
	if !p.IsSetSuccess() {
		return SnowWantedResult_Success_DEFAULT
	}
	return p.Success
}
func (p *SnowWantedResult) SetSuccess(x interface{}) {
	p.Success = x.(*SnowResponse)
}

var fieldIDToName_SnowWantedResult = map[int16]string{
	0: "success",
}

func (p *SnowWantedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SnowWantedResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_SnowWantedResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *SnowWantedResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = NewSnowResponse()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *SnowWantedResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("wanted_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *SnowWantedResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *SnowWantedResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SnowWantedResult(%+v)", *p)
}

func (p *SnowWantedResult) DeepEqual(ano *SnowWantedResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *SnowWantedResult) Field0DeepEqual(src *SnowResponse) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}
