using System.Net;
using Gore;
using Microsoft.AspNetCore.Server.Kestrel.Core;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

builder.Services.AddCors(o => {
    o.AddPolicy("dev", p => {
        // p.AllowAnyOrigin();
        //AllowAnyHeader().AllowAnyMethod().AllowCredentials()
    });
});

builder.Services.AddGrpc(o => o.EnableDetailedErrors = true);

builder.WebHost.ConfigureKestrel(o => {
    o.Listen(IPAddress.Any, 5006, lo => {
        lo.Protocols = HttpProtocols.Http2; 
        lo.UseHttps(@"C:\.x\.dmp\cert\localdev.pfx", "password");
    });
});

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseCors("dev");
// app.UseHttpsRedirection();

app.MapGrpcService<DataTransferProcedure>();

var summaries = new[]
{
    "Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering", "Scorching"
};

app.MapGet("/weatherforecast", () =>
{
    var forecast =  Enumerable.Range(1, 5).Select(index =>
        new WeatherForecast
        (
            DateOnly.FromDateTime(DateTime.Now.AddDays(index)),
            Random.Shared.Next(-20, 55),
            summaries[Random.Shared.Next(summaries.Length)]
        ))
        .ToArray();
    return forecast;
})
.WithName("GetWeatherForecast")
.WithOpenApi();

app.Run();

record WeatherForecast(DateOnly Date, int TemperatureC, string? Summary)
{
    public int TemperatureF => 32 + (int)(TemperatureC / 0.5556);
}
