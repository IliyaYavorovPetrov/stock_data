package com.stock.server.config;

import com.datastax.oss.driver.api.core.CqlSession;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class CassandraConfig {
    @Bean
    public CqlSession toCqlSession() {
        return CqlSession.builder().build();
    }
}
