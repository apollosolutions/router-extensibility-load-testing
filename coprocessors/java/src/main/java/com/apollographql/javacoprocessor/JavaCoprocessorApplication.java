package com.apollographql.javacoprocessor;

import java.util.ArrayList;
import java.util.Collections;
import java.util.UUID;
import java.util.LinkedHashMap;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonValue;

import io.jsonwebtoken.Jwts;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
@SpringBootApplication
public class JavaCoprocessorApplication {
    final String JWT_SECRET = System.getenv("JWT_SECRET") != null ? System.getenv("JWT_SECRET") : "apollo";

    RouterPayload sendUnauthenticated(RouterPayload payload) {
        payload.setControl(new CoprocessorControl(401));
        return payload;
    }

    @PostMapping("/client-awareness")
    RouterPayload handleClientAwareness(@RequestBody RouterPayload payload) {
        if (payload.getStage() != CoprocessorStage.ROUTER_REQUEST) {
            return payload;
        }

        var headers = payload.getHeaders();
        if (headers == null || headers.get("authentication") == null) {
            return sendUnauthenticated(payload);
        }

        var tokenString = headers.get("authentication").get(0).split("Bearer ")[1];
        if (tokenString == null) {
            return sendUnauthenticated(payload);
        }

        var keyBytes = JWT_SECRET.getBytes();
        if (keyBytes.length < 32) {
            var temp = new byte[32];
            System.arraycopy(keyBytes, 0, temp, 0, keyBytes.length);
            keyBytes = temp;
            temp = null;
        }

        try {
            var token = Jwts.parserBuilder().setSigningKey(keyBytes).build().parseClaimsJws(tokenString);
            var claims = token.getBody();

            var clientName = new ArrayList<String>();
            clientName.add(claims.get("client_name") != null ? claims.get("client_name").toString() : "coprocessor");
            var clientVersion = new ArrayList<String>();
            clientVersion
                    .add(claims.get("client_version") != null ? claims.get("client_version").toString() : "loadtest");
            headers.put("apollographql-client-name", clientName);
            headers.put("apollographql-client-version", clientVersion);
        } catch (Exception ex) {
            return sendUnauthenticated(payload);
        }

        return payload;
    }

    @PostMapping("/guid-response")
    RouterPayload handleGuidResponse(@RequestBody RouterPayload payload) {
        if (payload.getStage() != CoprocessorStage.ROUTER_RESPONSE) {
            return payload;
        }

        var headers = payload.getHeaders();
        var values = new ArrayList<String>();
        for (int i = 0; i < 10; i++) {
            values.add(i, UUID.randomUUID().toString());
        }

        headers.put("GUID", values);
        payload.setHeaders(headers);

        return payload;
    }

    @PostMapping("/static-subgraph")
    RouterPayload handleStaticSubgraph(@RequestBody RouterPayload payload) {
        if (payload.getStage() != CoprocessorStage.SUBGRAPH_RESPONSE) {
            return payload;
        }

        var headers = payload.getHeaders();
        var values = new ArrayList<String>();
        values.add(0, "coprocessor");
        headers.put("source", values);
        payload.setHeaders(headers);

        return payload;
    }

    public static void main(String[] args) {
        var port = System.getenv("PORT");
        if (port == null) {
            port = "3000";
        }

        var app = new SpringApplication(JavaCoprocessorApplication.class);
        app.setDefaultProperties(Collections.singletonMap("server.port", port));
        app.run(args);
    }
}

@JsonInclude(JsonInclude.Include.NON_ABSENT)
class CoprocessorControl {
    private Integer breakVal;

    CoprocessorControl() {
    }

    CoprocessorControl(Integer breakVal) {
        this.breakVal = breakVal;
    }

    public Integer getBreak() {
        return this.breakVal;
    }

    public void setBreak(Integer breakVal) {
        this.breakVal = breakVal;
    }
}

enum CoprocessorStage {
    ROUTER_REQUEST("RouterRequest"),
    ROUTER_RESPONSE("RouterResponse"),
    SUBGRAPH_REQUEST("SubgraphRequest"),
    SUBGRAPH_RESPONSE("SubgraphResponse");

    private String name;

    private CoprocessorStage(String name) {
        this.name = name;
    }

    @JsonValue
    public String getName() {
        return this.name;
    }

    public void setName(String name) {
        this.name = name;
    }
}

@JsonInclude(JsonInclude.Include.NON_ABSENT)
class RouterPayload {
    private Object control;
    private LinkedHashMap<String, ArrayList<String>> headers;
    private String id;
    private String method;
    private String sdl;
    private CoprocessorStage stage;
    private Integer version;

    RouterPayload() {
        this.control = "continue";
        this.id = "";
    }

    public Object getControl() {
        return this.control;
    }

    public void setControl(Object control) {
        this.control = control;
    }

    public LinkedHashMap<String, ArrayList<String>> getHeaders() {
        return this.headers;
    }

    public void setHeaders(LinkedHashMap<String, ArrayList<String>> headers) {
        this.headers = headers;
    }

    public String getId() {
        return this.id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getMethod() {
        return this.method;
    }

    public void setMethod(String method) {
        this.method = method;
    }

    public String getSDL() {
        return this.sdl;
    }

    public void setSDL(String sdl) {
        this.sdl = sdl;
    }

    public CoprocessorStage getStage() {
        return this.stage;
    }

    public void setStage(CoprocessorStage stage) {
        this.stage = stage;
    }

    public Integer getVersion() {
        return this.version;
    }

    public void setVersion(Integer version) {
        this.version = version;
    }
}
