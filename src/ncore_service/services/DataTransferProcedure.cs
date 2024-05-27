using Grpc.Core;
using Proto;

namespace Gore;

public class DataTransferProcedure : DataTransferService.DataTransferServiceBase
{
    public override Task<DataTransferPaystub> TransferData(DataTransferPayload request, ServerCallContext context)
    {
        Console.WriteLine(context);
        Console.WriteLine(request.IsFine);
        Console.WriteLine(request.Data);
        Console.WriteLine(request.Action);

        if (request.IsFine == false)
        {   
            return Task.FromResult(new DataTransferPaystub { Ack = false });
        }

        // return base.TransferData(request, context);
        return Task.FromResult(new DataTransferPaystub { Ack = true });
    }
}
