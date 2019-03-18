package xdean.mini.boardgame.server.model.entity;

import javax.annotation.CheckForNull;

import com.fasterxml.jackson.annotation.JsonIgnore;

import io.swagger.annotations.ApiModelProperty;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.Builder.Default;

@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder(toBuilder = true)
public class UserProfileEntity {

  @JsonIgnore
  int id;

  @Default
  @ApiModelProperty(position = 0, example = "dean")
  String nickname = "";

  @CheckForNull
  @ApiModelProperty(position = 1, example = "true")
  Boolean male;

  @Default
  @ApiModelProperty(position = 2, example = "http://cdn.aixifan.com/dotnet/20130418/umeditor/dialogs/emotion/images/ac3/01.gif")
  String avatarUrl = "";
}
