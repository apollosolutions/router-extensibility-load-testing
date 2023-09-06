using System.IdentityModel.Tokens.Jwt;
using System.Security.Cryptography;
using System.Text;
using System.Text.Json.Serialization;
using Microsoft.IdentityModel.Tokens;

var JWT_SECRET = Environment.GetEnvironmentVariable("JWT_SECRET") ?? "apollo";

var builder = WebApplication.CreateBuilder(args);
builder.Services.AddControllers();

var getRequestPayload = (CoprocessorRequest request) =>
{
    var payload = request;
    if (payload.Headers == null)
    {
        payload.Headers = new Dictionary<string, string[]>();
    }

    return payload;
};

var sendUnauthenticated = (CoprocessorRequest payload) =>
{
    payload.Control = new CoprocessorControl
    {
        Break = 401,
    };

    return payload;
};

var handleClientAwareness = (CoprocessorRequest request) =>
{
    var payload = getRequestPayload(request);
    if (payload.Stage != CoprocessorStage.RouterRequest)
    {
        return payload;
    }

    if (payload.Headers == null || !payload.Headers.ContainsKey("authentication"))
    {
        return sendUnauthenticated(payload);
    }

    var tokenString = payload.Headers["authentication"][0].Split("Bearer ").Last();
    if (tokenString == null || tokenString.StartsWith("Bearer"))
    {
        return sendUnauthenticated(payload);
    }

    var token = new JwtSecurityToken(tokenString);

    var keyBytes = Encoding.UTF8.GetBytes(JWT_SECRET);
    if (keyBytes.Length < 32 && token.SignatureAlgorithm == "HS256")
    {
        // Thanks .NET
        Array.Resize(ref keyBytes, 32);
    }

    var hmac = new HMACSHA256(keyBytes);
    var validation = new TokenValidationParameters
    {
        IncludeTokenOnFailedValidation = true,
        IssuerSigningKey = new SymmetricSecurityKey(hmac.Key),
        RequireAudience = false,
        RequireExpirationTime = false,
        RequireSignedTokens = false,
        ValidateAudience = false,
        ValidateIssuer = false,
        ValidateLifetime = false,
    };
    var handler = new JwtSecurityTokenHandler();

    try
    {
        handler.ValidateToken(tokenString, validation, out var validatedToken);
    }
    catch (Exception)
    {
        return sendUnauthenticated(payload);
    }

    var claims = token.Claims.ToList();

    payload.Headers["apollographql-client-name"] = new string[] { claims.FirstOrDefault(c => c.Type == "client_name")?.Value ?? "coprocessor" };
    payload.Headers["apollographql-client-version"] = new string[] { claims.FirstOrDefault(c => c.Type == "client_version")?.Value ?? "loadtest" };

    return payload;
};

var handleGuidResponse = (CoprocessorRequest request) =>
{
    var payload = getRequestPayload(request);
    if (payload.Stage != CoprocessorStage.RouterResponse)
    {
        return payload;
    }

    payload.Headers!["GUID"] = new string[10];
    for (int i = 0; i < payload.Headers["GUID"].Length; i++)
    {
        payload.Headers["GUID"][i] = Guid.NewGuid().ToString();
    }

    return payload;
};

var handleStaticSubgraph = (CoprocessorRequest request) =>
{
    var payload = getRequestPayload(request);
    if (payload.Stage != CoprocessorStage.SubgraphResponse)
    {
        return payload;
    }

    payload.Headers!["source"] = new string[] { "coprocessor" };

    return payload;
};

var port = Environment.GetEnvironmentVariable("PORT") ?? "3000";
var app = builder.Build();
app.UseAuthorization();
app.MapPost("/client-awareness", handleClientAwareness);
app.MapPost("/guid-response", handleGuidResponse);
app.MapPost("/static-subgraph", handleStaticSubgraph);
Console.WriteLine($"Starting on port {port}");
app.Run($"http://0.0.0.0:{port}");

class CoprocessorBody
{
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
    public dynamic? Data { get; set; }
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
    public dynamic? Errors { get; set; }
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
    String? Query { get; set; }
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
    String? OperationName { get; set; }
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
    public dynamic? Variables { get; set; }
}

class CoprocessorContext
{
    public dynamic? Entries { get; set; }
}

class CoprocessorControl
{
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
    public int? Break { get; set; }
}

[JsonConverter(typeof(JsonStringEnumConverter))]
enum CoprocessorStage
{
    RouterRequest,
    RouterResponse,
    SubgraphRequest,
    SubgraphResponse,
}

class CoprocessorRequest
{
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
    public CoprocessorBody? Body { get; set; }
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
    public CoprocessorContext? Context { get; set; }
    public dynamic Control { get; set; } = "continue";
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
    public Dictionary<string, string[]>? Headers { get; set; }
    public String Id { get; set; } = "";
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
    public String? Method { get; set; }
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
    public String? SDL { get; set; }
    public CoprocessorStage Stage { get; set; }
    public int Version { get; set; }

    // router* stage specific
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
    public String? Path { get; set; }

    // router response stage specific
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
    public int? StatusCode { get; set; }

    // subgraph* stage specific
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
    public String? ServiceName { get; set; }
    [JsonIgnore(Condition = JsonIgnoreCondition.WhenWritingNull)]
    public String? URI { get; set; }
}
