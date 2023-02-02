package com.stock.server.config;

import org.flywaydb.core.Flyway;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.cassandra.core.CassandraOperations;

@Configuration
public class CassandraFlywayConfiguration {
    @Bean
    public Flyway flyway(CassandraOperations cassandraOperations) {
        Flyway flyway = Flyway.configure()
                .dataSource("jdbc:cassandra://localhost:9042/stocks", "cassandra", null)
                .locations("classpath:/db/cassandra/migration")
                .load();
        flyway.migrate();
        return flyway;
    }
}
