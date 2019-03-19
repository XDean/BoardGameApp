package xdean.mini.boardgame.server.model.param;

import com.fasterxml.jackson.databind.annotation.JsonSerialize;

import lombok.Builder;
import lombok.Value;

@Value
@Builder
@JsonSerialize
public class ExitGameResponse {
}
